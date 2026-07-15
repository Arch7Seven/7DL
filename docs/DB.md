## Tables :
- users
- transfers
- settings

### Users
- id (PK)           (int, nonnull)
- uuid              (string, uuid, nonnull)
- username          (string, nonnull)
- created_at        (int, timestamp, nonnul)
- active            (bool, nonnul)
- settings_id       (FK)

### Credentials
- uuid              (FK, PK)
- password          (string, hash, local only)

### Settings
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



