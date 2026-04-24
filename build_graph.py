import json
import sys
from pathlib import Path
from graphify.detect import detect
from graphify.extract import collect_files, extract
from graphify.build import build_from_json
from graphify.cluster import cluster as cluster_graph
from graphify.analyze import god_nodes, surprising_connections
from graphify.report import generate
from graphify.export import to_html
import networkx as nx
from networkx.readwrite import json_graph

# Run from the directory passed as argument, or current directory
target = Path(sys.argv[1]) if len(sys.argv) > 1 else Path('.')

print("="*60)
print(f"Building graphify knowledge graph for: {target.resolve()}")
print("="*60)

out_dir = target / 'graphify-out'
out_dir.mkdir(exist_ok=True)

# Step 1: Detect files
detect_data = detect(target)
(out_dir / '.graphify_detect.json').write_text(json.dumps(detect_data, indent=2))
print(f"\n✓ Detected: {detect_data['total_files']} files (~{detect_data['total_words']} words)")
for ftype, files in detect_data.get('files', {}).items():
    if files:
        print(f"    {ftype}: {len(files)} files")

if detect_data['total_files'] == 0:
    print("No supported files found. Exiting.")
    sys.exit(1)

# Step 2: AST extraction
code_files = []
for f in detect_data.get('files', {}).get('code', []):
    p = Path(f)
    code_files.extend(collect_files(p) if p.is_dir() else [p])

if code_files:
    ast_data = extract(code_files)
    (out_dir / '.graphify_ast.json').write_text(json.dumps(ast_data, indent=2))
    print(f"✓ AST: {len(ast_data['nodes'])} nodes, {len(ast_data['edges'])} edges")
else:
    ast_data = {'nodes': [], 'edges': [], 'input_tokens': 0, 'output_tokens': 0}
    (out_dir / '.graphify_ast.json').write_text(json.dumps(ast_data, indent=2))
    print("✓ AST: No code files — skipped")

# Step 3: Build graph
extraction = {
    'nodes': ast_data.get('nodes', []),
    'edges': ast_data.get('edges', []),
    'hyperedges': ast_data.get('hyperedges', []),
    'input_tokens': ast_data.get('input_tokens', 0),
    'output_tokens': ast_data.get('output_tokens', 0),
    'total_files': detect_data.get('total_files', 0),
    'total_words': detect_data.get('total_words', 0),
}

G = build_from_json(extraction)
print(f"✓ Graph: {G.number_of_nodes()} nodes, {G.number_of_edges()} edges")

# Cluster
communities = cluster_graph(G)
print(f"✓ Clustering: {len(communities)} communities detected")

# Analyze
gods = god_nodes(G)
surprises = surprising_connections(G, communities)
print(f"✓ Analysis: {len(gods)} god nodes, {len(surprises)} surprising connections")

# Export JSON
graph_data = json_graph.node_link_data(G)
(out_dir / 'graph.json').write_text(json.dumps(graph_data, indent=2))
print("✓ Exported: graphify-out/graph.json")

# Generate report
analysis = {
    'communities': {str(k): v for k, v in communities.items()},
    'cohesion': {},
    'god_nodes': gods,
    'surprises': surprises,
}
(out_dir / '.graphify_analysis.json').write_text(json.dumps(analysis, indent=2))

report = generate(G, communities, {}, {}, gods, surprises, extraction, {'input_tokens': 0, 'output_tokens': 0}, str(target))
(out_dir / 'GRAPH_REPORT.md').write_text(report, encoding='utf-8')
print("✓ Exported: graphify-out/GRAPH_REPORT.md")

# Generate HTML
try:
    to_html(G, communities, str(out_dir / 'graph.html'))
    print("✓ Exported: graphify-out/graph.html")
except Exception as e:
    print(f"⚠ HTML export skipped: {e}")

print("\n" + "="*60)
print("✓ graphify complete")
print("="*60)
print("\nOutputs:")
print("  • graphify-out/graph.json      — GraphRAG-ready, queryable")
print("  • graphify-out/graph.html      — interactive visualization")
print("  • graphify-out/GRAPH_REPORT.md — architecture summary")
