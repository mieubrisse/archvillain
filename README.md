Archvillain: Manage Your Minions
================================
Like any self-respecting evil genius, you're planning world domination.

Only, world domination requires so much _stuff_ - publishing on social media, booking flights, boilerplate code, responding to Aunt Mary's text.

What a drag. Life's not all doom lasers and fun.

But wait... evil geniuses have minions! Agentic AIs to help you get back to the fun parts.

Only, minions are kinda dumb. They often get confused, and do silly things like [delete your entire computer](https://forum.cursor.com/t/cursor-yolo-deleted-everything-in-my-computer/103131/3).

If only you had a command center to manage your minions, letting them go when they're doing good work and corralling them when they're going off track...

This is Archvillain.

### In plain English?
Archvillain is a platform for amplifying your personal productivity by using AI agents to complete tasks.

It lets you launch AI agents in safe sandboxes, monitor their progress, step in when necessary, save their output, and repeat the cycle until the task is done.

It is designed around the philosophy that you should spend your time doing [unique work that only you can do](https://mieubrisse.substack.com/p/the-goal-is-unique-work), and the AI should do the rest.

It is best for tasks with a start and end, like:

- Building a feature in code
- Writing a blog post
- Generating images for the blog post
- Booking flights
- Renewing a driver's license online

It is not for launching and running agents that run forever.

### Why Archvillain?
AI is great. But it has some flaws:

- In non-YOLO mode, it's too slow and unsure
- In YOLO mode, it sometimes does insane things [like delete your entire filesystem](https://forum.cursor.com/t/cursor-yolo-deleted-everything-in-my-computer/103131/3)
- Without standardized guidance, its output is too wide
- Sometimes it's stuck producing junk and we need to backtrack

As a community, we're learning to build around this:

- Commit early, and often
- Write specs before building features (e.g. [vibe speccing](https://lukebechtel.com/blog/vibe-speccing))
- Run AIs inside sandboxed containers (e.g. [using Claude Code in devcontainers](https://timsh.org/claude-inside-docker/); [container-use](https://github.com/dagger/container-use))
- Build standardized guidance given to all AI tools (e.g. [treating prompts as code](https://mariozechner.at/posts/2025-06-02-prompts-are-code/); [inputs, not outputs](inputs not outputs))

Archvillain incorporates all these lessons, plus some innovations of its own, into a single open-source tool.

### An Archvillain example
It's easiest to see by example. 

Each outcome you want to accomplish in Archvillain is called a task.

Let's say that our task is to generate a cover image for a blog post we've written.

We'll create a new task in Archvillain, tell it that "We want to generate an image for the attached blog post", and attach the blog post text as a PDF. We can optionally pick the LLM that will process our task if we like. In this case we'll use GPT-4 for its image generation capabilities.

In Archvillain's GUI we see a new tab get created for our task.

We click inside the tab and see a directed acyclic graph (DAG) representing the work that's been started.

The root node (Node 1) represents a workspace containing files with the initial context that we provided: the text stating the goal of the task, and the PDF containing the blog post text.

We then see an edge being created off Node 1. This edge represents the AI's work. The AI's work is being done inside a Docker container, with the files from Node 1 mounted into the container. Because it's running inside a container sandbox, the AI can run without restrictions (e.g. if the AI is Claude, we can approve all tools).

There's a little error icon on the edge, because the GPT-4 running inside the container needs our OpenAI credentials. We click on it, and 1Password pops up, prompting us to authorize the use of our OpenAI credentials in the container. We authorize, and the edge starts starts flashing as it works.

Clicking on the edge while it's in progress shows GPT-4's thinking output, and the list of changes the GPT-4 is making to the files inside the container.

At any point while the AI is working, we can SSH inside the container and explore what the AI is doing.

When GPT-4 finishes its work, the edge off Node 1 stops flashing and a new node (Node 2) is created at the unfinished end of the edge.

Clicking on Node 2 shows the contents of the workspace after GPT-4's work: a new `image.png` file is now in the workspace, next to our initial context and PDF containing the blog post text.

At this point, we have several options depending on what we think of the output:

1. We might like the output `image.png` in Node 2 and decide the task is complete. We can then download any files we want from Node 2 - just `image.png` if we please, or everything.
1. We might not like the output `image.png` and want to refine it with further instructions. In this case, we can click "+" on Node 2, add more context, and fire off more AI work which will be represented as an edge off Node 2 which will eventually result in a new node.
1. We might not like the output image and decide we want GPT-4 to retry from our initial prompt. In this case, we can click "+" on Node 1 and fire off another attempt using the context in Node 1. The result will be another edge off Node 1 - a parallel line of exploration independent of the work that resulted in Node 2.
1. We might not like the output image and decide we want to update the initial context that we gave in Node 1. In this case, we'll choose "Duplicate" on Node 1 to get ready to create a new root node (Node 1b), update the files/context in it, and then get ready to fire off work from Node 1b that's independent and parallel to anything off of Node 1.
1. We could decide that we want to update the guidance we give to _all_ agents in the future. The Archvillain platform provides a Git repository of Markdown files that you can manage, nad optionally include in every agent run. You can specify detailed rules for when and how these Markdown files get included (e.g. only include certain files if the agent is running a coding task, or only include certain files if the agent is Claude but not GPT).

Because every node in Archvillain is a workspace containing files containing context, we can explore different lines of thinking in parallel without fear of losing our work. We can even combine files from different lines of thinking (e.g. create a new node that uses `image2.png` from Node X and `image3.png` from Node Y), even pulling files from nodes from other tasks if you need.

In this way, Archvillain formalizes the "human gives context then AI works until it comes back" workflow that currently exists with tools like ChatGPT and Claude. 

Archvillain improves on this workflow by keeping state at each step of the conversation, allowing you to explore freely or even put down the conversation and come back days later when it's convenient to you.

You can even kick off multiple pieces of work in parallel (multiple edges, running at the same time).

### Archvillain ideas
Archvillain supports any type of AI work you please. Just provide a Docker image with the tools you need and let the AI use them in the container. Some other ideas:

- **Making changes to a code repo:** You'd start the task by providing the code repository, allowing the AI to then make changes. When you're happy with the work, you'd commit the files of the node you're happy with. To help with this workflow, Archvillain provides built-in tools to start a task from a commit in a Git repo, and to commit the files in a node back to the Git repo.
- **Booking flights:** You can have agents inside of Archvillain do work on the web. Simply run a container that has a browser, point the agent running in the container to the browser, and give it instructions. Groo provides a utility Docker image that has a headless browser for just this purpose. Secrets like site login credentials and credit card information can be given to the container on an on-demand basis.

