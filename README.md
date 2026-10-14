# Plus ou moins

Jeu du plus ou moins en ligne de commande. Le programme choisit un nombre
secret entre 1 et 100, le joueur propose des nombres et le programme répond
« C'est plus ! » ou « C'est moins ! » jusqu'à ce que le nombre soit trouvé.
Le meilleur score (le moins d'essais) est conservé dans `record.txt`.

## Installation

Prérequis : Go installé (version indiquée dans `plusmoins/go.mod`).

```
git clone https://github.com/geekworldcode/plusmoins.git
cd plusmoins/plusmoins
go build -o plusmoins.exe ./cmd/plusmoins
```

## Utilisation

Toutes les commandes se lancent depuis le dossier `plusmoins/` (celui qui
contient `go.mod`).

```
go run ./cmd/plusmoins
go run ./cmd/plusmoins -max 50
.\plusmoins.exe -max 50
```

| Option    | Défaut       | Rôle                                        |
|-----------|--------------|---------------------------------------------|
| `-max`    | `100`        | Borne haute du nombre secret                |
| `-secret` | `0` (hasard) | Impose le nombre secret, pour les tests     |

Aide : `go run ./cmd/plusmoins -h`

Pendant la partie :

- `q`, `Q` ou `quitter` (quelle que soit la casse) quitte la partie ;
- une saisie vide, un texte qui n'est pas un nombre ou un nombre hors de
  `1-max` est refusé et le nombre est redemandé ;
- le record est affiché au lancement et enregistré dans `record.txt` quand
  il est battu.

## Tests

Tests automatiques :

```
go test ./...
```

Test manuel avec les saisies pièges. La sortie doit être identique à
`testdata/attendu.txt` (supprimer `record.txt` avant, pour partir du même
état) :

```
# Windows (PowerShell)
Get-Content testdata\pieges.txt | go run ./cmd/plusmoins -secret 42
# Linux, Mac
cat testdata/pieges.txt | go run ./cmd/plusmoins -secret 42
```

## Structure

```
plusmoins/
├── go.mod
├── cmd/plusmoins/        main.go : assemble les morceaux
├── internal/
│   ├── game/             règles et état (secret, essais)
│   ├── parser/           analyse des saisies (ParseGuess)
│   ├── score/            record sur disque (record.txt)
│   └── ui/console/       entrées et affichage : ask, welcome, showresult,
│                         showerror, showrecord, showfatal, color
└── testdata/             saisies pièges et sortie attendue
```
