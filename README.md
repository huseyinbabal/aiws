You don't need to memorize AWS CLI commands anymore. Use aiws, AI-driven AWS CLI
that uses Chat GPT to generate AWS CLI commands from your prompts.

## Prerequisites
- [AWS CLI](https://aws.amazon.com/cli/)
- [OPEN AI API KEY](https://platform.openai.com/account/api-keys)

## Installation
### MacOS
```bash
brew tap huseyinbabal/tap
brew install aiws
```
### Linux & Windows
Please check [release pages](https://github.com/huseyinbabal/aiws/releases) to see linux and windows packages.

## Configuration
```bash
export OPEN_AI_API_KEY=<open_api_key>
```

## Examples

> **Note**
> More prompts will be here soon

### List EC2 instances in us-west-2
```shell
aiws list ec2 in us-west-2
```
![aiws-list-ec2](https://user-images.githubusercontent.com/1237982/228968604-953dd131-ba9d-47dd-a8d0-19bf59e37664.gif)
