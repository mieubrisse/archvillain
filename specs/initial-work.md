Your job is to build out a v0 of the grand vision described in this project's README.md.

Workflow
--------
For this MVP, we're shooting to enable a workflow of allowing a user to generate an image for a blog post using OpenAI's GPT, and iterate on the image as they please.

The user's workflow should be:

1. Access the Archvillain web server
1. The user gets presented with a web app that looks like a sidebar (with nothing in it except a "New Task" button), and a main pane that's empty
1. Press "New Task"
1. Get presented a context dialog explaining the user's intitial prompt
1. User presses "Go", which triggers:
    1. An entry for the task appears in the left sidebar is created, selected
    1. The main display for the selected task shows an infinite canvas graph displaying a directed acyclic graph (DAG) with a single root node (Node 1), and a single edge that's flashing as the AI works
1. Clicking on Node 1 shows the context that the user provided: the prompt that they started things with
1. Clicking on the edge shows details about the AI's execution:
    - The Docker image that is being executed
    - The output logs as the AI thinks
1. When the AI is done thinking, a new node (Node 2) is created at the end of the edge. Clicking on Node 2 shows the same files as were inputted, only with an additional `image.png` file representing the generated output.
1. While Node 2 is selected, selecting `image.png` shows the contents of the image file inside the Archvillain web app
1. There is a "download" button next to the `image.png` file, and clicking it downloads the image

Architecture
------------
- The code should be written as a Typescript React app with Node backend, created with Vite.
- There should be a devcontainer for running the Archvillain server app
- The devcontainer should mount the Docker socket from the external machine, so that the Archvillain server running in the devcontainer can manipulate Docker
- When the user creates a new task, a new volume is created that 
- All containers and volumes created by Archvillain should get a label called "archvillain.version" with value "1.2.3" (which is how we'll find these later)

TODO workspace directory

TODO how to store node data on the server

TODO use the MUI React component framework

1. There should be a server and a CLI
1. The sever 




Style
-----
It is essential that you write well-factored code, with small functions and good decomposition. This helps keep the code modular, which allows ease of future editing.
