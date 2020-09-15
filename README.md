Uses the YouTube Data Web API to check the recent activities of channels to display a list 
of new videos.

### Configuration in config/config.go

Modify `seenFileName` to a path you like, set your YouTube key in `youTubeKey` (like `AIzaSy...`)
and change the returned string array of the `GetYouTubeChannelsToCheck()` function.

To extract the Channel ID from channels like this https://www.youtube.com/user/vsauce where the 
URL does not contain the ID,
this website can be used: https://commentpicker.com/youtube-channel-id.php

---

The file specified by `seenFileName` contains the seen videos in this format:

```bash
Medical Medium|The Medical Medium Answering Your Food & Health Questions
Medical Medium|Why Celery Juice Is Healing Millions
...
```

### Output

Running this produces an output like this and only lists new videos:

```bash
Medical Medium   https://www.youtube.com/channel/UCUORv_qpgmg8N5plVqlYjXg
  - The Medical Medium Answering Your Food & Health Questions
  - Why Celery Juice Is Healing Millions
```

It uses the https://www.googleapis.com/youtube/v3/activities API call which returns just the last 5 videos, but 
more are really not needed.

As a note, running this fresh always displays all videos as new because `seen-youtube.txt` is not filled yet.
