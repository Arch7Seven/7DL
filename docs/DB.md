## Tables :
- users
- transfers
- settings

### Identities
- uuid                      (uuid, PK)
- username                  (string, nonnul)
- email                     (string, nullable)
- password                  (string, hash, local only)

### Users
- uuid                      (PK, FK (Identities uuid))
- created_at                (int, timestamp, nonnul)
- display_name              (string)

### Settings
- uuid                      (PK, FK (Users uuid))
- preferred_video_codec     (string, nullable)
- preferred_audio_codec     (string, nullable)
- cookies                   (json, nullable)

#### Cookies:
un object json avec une key par site (youtube, instagram, etc)
{
    "youtube": "cookie",
    "instagram": "cookie"
}

peut être mettre les cookies dans des [] ?



