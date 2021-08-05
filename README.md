# ![logo](public/mnimidamon-frontend-icon.svg) mnimidamon frontend
Frontend graphical interface for mnimidamon application developed with Fyne.io toolkit.

# Http client code generation

Generate  client code from the swagger.yml file located at the commit tagged `v0.1` in mnimidamon-backend repository:

`swagger generate client -f https://raw.githubusercontent.com/mnimidamon/mnimidamon-backend/0ac9fea6e0af1b8b14275fc851a1e1f9cb1af486/public/spec/swagger.yaml -A mnimidamon`