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
