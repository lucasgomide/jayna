# Jayna

> Jayna: Shape of a spider

Transforms everything into anything. Just say what you wanna transforms and.. voilà

## Inputs supporteds
- CSV

# Output supporteds
- JSON

## Usage:

```bash
jayna -d "header1;header2;header3;\ncontent1;content2;1" -o json
jayna -f file.csv -o json
```
Output expected
```json
{header1: "content1", header2: "content2", header3: 1}
```
