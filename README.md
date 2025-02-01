# TUISpeak

This is a sample project to help me learn [BubbleTea](https://github.com/charmbracelet/bubbletea).

## Requirements
Ubuntu 22.04+

Packages TBD.  I know it depends on `speech-dispatcher`, `espeak`, and `festival`, but I'm not 100% sure.

## Why?
Sometimes, life throws you curveballs and you lose your voice.

This is for those times.

# Sample Configuration
At the moment, as part of this proof of concept, `boards.yml` is the only supported configuration.  This will improve in time via [Viper](https://github.com/spf13/viper) integration.

``` yaml
boards:
  - title: Everyday Qs
    questions:
      - May I please be suctioned?
      - Please refill my water.
      - May I please have a breathing treatment?
      - Thank you.

  - title: Personable Qs
    questions:
      - It's good to see you.  How are you doing?
      - What's new?
```

# Usage
Use `j`/`k` or `Up`/`Down` to select your pre-defined choices, then `s` or `Enter` to have it read back to you.

Use `h`/`l` or `Left`/`Right` to move between boards, assuming you have more than one.

## Screenshot
<img src="screenshot.png" width="720" height="480" alt="Screenshot">

# Licensing

This project is dual-licensed under the MIT License for open-source use and a commercial license for proprietary use. See the [LICENSE](LICENSE) file for more details.
