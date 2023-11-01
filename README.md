## ripcurl

Easily snag text from a web page and use it however you like. The advantage
of this tool over other tools and methods, is that ripcurl can bypass basic
Cloudflare checks and does a far better job of "cleaning" text of garbage input.
This tool can also easily be used in pipes or as a text filter for HTML.

---

Usage:

```sh
  ./ripcurl --url <url>
```

```sh
  cat out.html | ./ripcurl
  # or
  curl -fsSl -A "user_agent" | ./ripcurl
```

```sh
  ./ripcurl --url "<url>" > output.txt
  ./ripcurl --url <url> | less
```

Output:

```
Light Novel Volume 4 Episode 25

Sa Hyokyung was ambitious.
Like all men, he dreamed of reigning under the heavens. After all, he dreamed of living once, the pinnacle of Jianghu.
However, contrary to his ambition, reality was harsh.
In Jianghu, the Two Factions, Three Clans, Three Packs, and Three Manors were firmly established, and the Three Saints were walking around.
There was no place for the Seven Stars to intervene.
Although they managed to establish themselves in Hunan and gain fame, the limitations of the Seven Stars were clear.
Each member of the Seven Stars was clearly a master, but they did not reach the level where they could command Jianghu.
Although Sa Hyokyung possessed the force that overwhelms the other six, he was not unique enough like the Three Saints.
...
```

### A note on Text-to-speech
There is no easy way to configure tts across all platforms, there is no unified standard anywhere. This is an issue especially
with the intent for this tool to be used across Linux, Mac, Windows and Android. So I've so far decided to leave that up to you
to configure.

- On Linux I suggest using speech-dispatcher, alongside Piper-tts, or a wine prefix running a Windows TTS tool (like IVONA for example) with the `balcon.exe` CLI tool.
- On android I suggest using termux, with the termux API to handle native phone TTS.
- On windows, use `balcon.exe` alongside a TTS application.
- Idk about Mac or openBSD, sorry but I can assume that these above solutions will be the same.
- another "catch-all" solution is to use edge-tts or gtts-cli for Edge TTS from microsoft and Google TTS respectively.

NOTE that I am still working on a good way to auto-configure commands. My first idea is using placeholders like
```sh
[tts-to-wav]
command = "cat {{file}} | balcon -n Amy -i -w {{placeholder}}.wav"
```
where this is a TOML config file and `{{file}}` and `{{placeholder}}` or placeholder terms that the CLI tool auto-fills to allow you to use any TTS system you prefer.
In this case `balcon` is a simple wrapper around a wine prefix that uses `balcon.exe` this could also be done natively with Piper-tts (and my provided model if you wish)

Another example:
```sh
[tts-text]
command = "echo "{{text}}" | piper-tts --model ~/models/amy.onnx --output_raw | aplay -r 22050 -c 1 -f S16_LE -t raw"
```

Termux example:
```sh
[tts-text]
command = "termux-text-to-speech -t $(echo {{text}} | perl -pe 's/[^[:ascii:]]//g')"
```
