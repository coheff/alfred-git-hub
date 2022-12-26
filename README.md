# alfred-git-hub
An Alfred GitHub workflow powered by Go. Search is currently limited to repositories (release v0.1.0) for now. Support for more [SearchTypes](https://docs.github.com/en/graphql/reference/enums#searchtype) will be added later.

# Download
Grab the the latest version from the [releases page](https://github.com/coheff/alfred-git-hub/releases/tag/v0.1.0). Double click workflow file to import into Alfred.

If running on macOS Catalina or later, you _**MUST**_ add Alfred to the list of security exceptions for running unsigned software. Step-by-step instructions are available on the awgo wiki [here](https://github.com/deanishe/awgo/wiki/Catalina).

You can also grant permissions to this workflow only. To do this, install the workflow and within the workflow tab right click then select `open in Finder`. Then right-click on "alfred-go-drive-search" and select `open with Terminal`. Agree to open it. It will run and complain about not being started by Alfred. Ignore and close the terminal. After that, the workflow will work (until an update of the executable).

# Prequisites
In order to use this workflow you _**MUST**_ create your own GitHub OAuth2 app and generate a `client_id` & `client_secret`. Step-by-step instructions are available [here](https://docs.github.com/en/developers/apps/building-oauth-apps/creating-an-oauth-app). Once completed, copy and paste the `client_id` & `client_secret` into their respective workflow environment variables:

<img width="1008" alt="Screenshot 2022-12-26 at 15 00 54" src="https://user-images.githubusercontent.com/11773454/209561665-c661efba-4716-453c-9ca5-f92b04588957.png">

# Usage
- Trigger a repository search using the keyword `gh` followed by a search query.
- Check out the GitHub documentaion for query syntax - here's a [cheatsheet](https://gist.github.com/bonniss/4f0de4f599708c5268134225dda003e0).

![Screen Recording 2022-12-26 at 15 12 09](https://user-images.githubusercontent.com/11773454/209565558-fc5d9661-562a-4d16-9093-5424591fc883.gif)

# License
Distributed under the MIT License. See [LICENSE](https://github.com/coheff/alfred-git-hub/blob/main/LICENSE) for more information.
