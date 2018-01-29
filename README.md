# Engineering Notebook

This is a basic scaffolding for an engineering notebook. When you create a new section, you get the following directory structure:

```
    section_name
        |__ notes/
        |__ data/
        |__ scripts/
        |__ README.md
```

## Usage

The simple folder structure helps you organize ongoing work on an engineering problem.

As you collect supporting data, put the files in the `data` folder. As you build small pieces of code, collect them in the `scripts` folder. Likewise, any note taking along the way can live in the `notes` folder.

Keep re-writing the top-level `README.md` file as you refine your understanding of the problem. Ultimately, the README evolves into the document you present to your co-workers or other stakeholders that includes a high-level overview of the problem, the steps you took in your investigation, and solution you are proposing.

## Setup

When you run `notebook` there are two flags you can optionally pass:

```
-notepath string
    Main notebook directory set by env var NOTEPATH (default "notebook")
-section string
    Name for new section of this notebook (default "new_section")
```

By default, `notebook` with look for an environment variable called `NOTEPATH` that defines your main engineering notebook directory, but you can override this with the `notepath` flag.

You will want to use the `section` flag to name the section you are creating in a meaningful way for the new project.

Run:

```
./notebook -section=new_project && cd $NOTEPATH/new_project && ls
```

And you will see your new section ready to be filled up with interesting data, scripts, and notes!

