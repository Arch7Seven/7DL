
# Sign up
- Le client envoie les informations de connections (username, password) au backend
- Le backend créer l'identité et la persiste
- Le backend créer une session et réponds avec au client
- Le client fetch les user data liée au UUID de l'identity
- Si première connections, le client affiche les interfaces pour remplir certains champs comme display name
- Le backend reçoit les nouvelles user data et les persistes

