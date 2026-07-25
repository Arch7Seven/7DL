
## Edit user
*Le client est déjà connecté et a une session ouverte en court*
- Le client envoie une requête PATCH sur la route user
- Le backend la reçois sur le handler associé
- Le handler appelle auth pour vérifier la session, auth retourne le uuid de l'identity correspondante
- Le handler passe l'uuid à user