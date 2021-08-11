# ![logo](public/mnimidamon-frontend-icon.svg) mnimidamon frontend
Frontend graphical interface for mnimidamon application developed with Fyne.io toolkit.

# Http client code generation

Generate  client code from the swagger.yml file located at the commit tagged `v0.1` in mnimidamon-backend repository:

`swagger generate client -f https://raw.githubusercontent.com/mnimidamon/mnimidamon-backend/8ea2ec9edc702236b2694293939541a3c1c52f7a/public/spec/swagger.yaml -A mnimidamon`