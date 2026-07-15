
# Backend :
Le backend suit un schéma type de usecase -> port -> implémentation
Le usecase est l'orchestrateur qui déclenche une action
Le port est une interface qui normalise les appels peut importe l'implémentation

Par exemple pour un job :
1. Le client (frontend) lance un job
2. Le server backend le reçoit et transmet au usecase via un handler
3. Le usecase va générer un uuid, etc puis push le job dans la queue selon une interface (queue via strategy pattern)
4. Le front fait du polling toute les x secondes pour avoir le status
5. Le backend recupère le status via le port

en local la queue est juste une table de job partagée

en distant la queue est redis et le worker est séparé, on utilise donc redis stream (stream jobs et status):
le backend publie le job sur {jobs} et écoute en async sur status
le worker push le status sur {status} quand il ya une mise à jour
le backend la récupère et met à jour sa table de job

# Server
handler : parse le contenu d'une requête vers un usecase
routes : construit les routes
router : associe une route à un handler



ex changer username :

1. le client demande à changer d'username
2. le handler appelle auth pour vérifier le token et extrait le UUID
3. si token valide le handler appelle le usecase user qui modifie et persiste l'état