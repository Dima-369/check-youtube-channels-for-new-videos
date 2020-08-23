Uses the YouTube Data Web API to check the recent activities of channels to display a list 
of new videos.

---

Seen videos are simply stored in the file named `seen-youtube.txt` with this format:

```bash
Medical Medium|The Medical Medium Answering Your Food & Health Questions
Medical Medium|Why Celery Juice Is Healing Millions
...
```

### Configuration

Modify `seenFileName` in `helpers/files.go` to a path you like and modify
the `youTubeChannels` string array in `main.go` to your YouTube channels.

To extract the Channel ID from pages like this https://www.youtube.com/user/vsauce, 
this website can be used: https://commentpicker.com/youtube-channel-id.php

### Output

Running this produces an output like this and only lists new videos:

```bash
Medical Medium   https://www.youtube.com/channel/UCUORv_qpgmg8N5plVqlYjXg
  - Why Celery Juice Is Healing Millions
```

The https://www.googleapis.com/youtube/v3/activities API call returns the last 5 videos and more are really not needed.

As a note, running this fresh always displays all videos as new because `seen-youtube.txt` is not filled yet.
