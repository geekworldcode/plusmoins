# Plus ou moins

Jeu du plus ou moins en ligne de commande. Le programme choisit un nombre
secret entre 1 et `max`, le joueur propose des nombres et le programme répond
« plus » ou « moins » jusqu'à ce qu'il soit trouvé. Le meilleur score (le moins
d'essais) est conservé dans `record.txt`.

## Installation

Prérequis : Go (voir `go.mod`).

```
go build -o plusmoins.exe ./cmd/plusmoins
```

## Utilisation

```
go run ./cmd/plusmoins
go run ./cmd/plusmoins -max 50
go run ./cmd/plusmoins -h
```

| Option    | Défaut       | Rôle                                              |
|-----------|--------------|---------------------------------------------------|
| `-max`    | `100`        | Borne haute du nombre secret                      |
| `-secret` | `0` (hasard) | Impose le nombre secret, pour les tests           |
| `-record` | `record.txt` | Fichier où est enregistré le record               |

Pendant la partie, taper `q`, `quit` ou `quitter` (quelle que soit la casse)
quitte le jeu. Une saisie vide, non numérique ou hors de `1-max` est refusée et
le nombre est redemandé sans compter d'essai.

## Tests

```
go test ./...
```

Test manuel avec les saisies pièges (la sortie doit être identique à
`testdata/attendu.txt`) :

```
# Windows (PowerShell)
Get-Content testdata\pieges.txt | go run ./cmd/plusmoins -secret 42 -record $env:TEMP\record.txt
# Linux, Mac
cat testdata/pieges.txt | go run ./cmd/plusmoins -secret 42 -record /tmp/record.txt
```

## Structure

```
plusmoins/
├── cmd/plusmoins/     assemble les morceaux (main)
├── internal/
│   ├── game/          règles et état
│   ├── parser/        analyse des saisies
│   ├── score/         record sur disque
│   └── ui/console/    entrées et affichage (un package par fonction, plus color)
├── testdata/          saisies pièges et sortie attendue (avec les codes couleur ANSI)
└── README.md
```
