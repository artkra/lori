# Proposed protocol for turn-like server **lori**

- Starts with BEGIN word ("+++idspispopd_____" - 18 symbols)
- Then goes receiver (hash, name, address or some unique identifier = 64 symbols)
- Then goes type of payload ("T" = text, "B" = bytes)
- Then goes payload length (max 64*1024)
- Then goes message body (max 64*1024 symbols)
- Then goes end token ("_____idspispopd+++" - 18 symbols)

