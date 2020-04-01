# IFTTT Trigger Action

Trigger IFTTT with your custom event and values.

## Usage

Go to your IFTTT account and enable [**IFTTT Maker Webhooks**](https://ifttt.com/maker_webhooks), and save your own **token**.

In your GitHub action yaml:

```yaml
- name: "Let's my IFTTT know this"
  uses: cixo/ifttt-trigger@v0.3
  with:
    key: "YOUR_OWN_TOKEN"
    event: "build_success"
    value1: "demo"
```

P.S. You may store your token as [`Secrets`](https://help.github.com/en/actions/configuring-and-managing-workflows/creating-and-storing-encrypted-secrets).

Good luck!
