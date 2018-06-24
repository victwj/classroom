# Jekyll

Jekyll is a static site generator that works well with GitHub Pages.

Learned to develop my [personal website](https://victorwj.github.io) and replace the
previous pure HTML/CSS/JS version.

## Sources

- [Jekyll Docs](https://jekyllrb.com/docs/home/)

## Notes

- Ruby is a programming language, and RubyGems is a package management tool for Ruby. A package
is called a *gem*. Jekyll is a Ruby gem.

- To install jekyll, Ruby and RubyGems must first be installed. 

- Best practice is to not install gems as root. Set up `$GEM_HOME` and add `$GEM_HOME/bin` to `$PATH`.

- [Bundler](https://bundler.io) is another important Ruby gem. It is a tool to help manage project
environments, and makes it easy to manage dependencies (i.e. other gems, their dependencies, their versions).

- Jekyll transforms text. Without a tool like Jekyll, devs need to write content in pure HTML. With Jekyll, we can
break this down to a more maintainable process. We first make templates for our website, in Liquid HTML, to define
what things should look like. Then we write our content in Markdown. Jekyll will connect the two together and put the
content in the templates.

## Setup

```bash
mkdir website
cd website

# Generate a Gemfile in current directory
# Gemfile is how Bundler knows what the required dependencies are, and how to install them
bundle init

# Add the jeykll gem to Gemfile
bundle add jekyll

# Install the dependencies specified by Gemfile
# Creates Gemfile.lock, which is a snapshot of all dependency resolution and their versions
# This is useful because we can see what worked last time, and keep consistency for the future
# E.g. running bundle install in the future might install different versions, and portentially break things
bundle install
```

## Practice

- Important read:
[Variables](https://jekyllrb.com/docs/variables/) and [Directory
Structure](https://jekyllrb.com/docs/structure/). 

- To serve the website to `localhost:4000` and build to directory `_site`, run `jekyll serve`.

- `_layouts/`

This is the directory where all the templates are. The underscore signifies
that this directory is not meant to be rendered for the final site, in this
case because it contains templates. We render content *using* templates, and usually
not just templates by themselves. For example, without the underscore, the `index.html` template
will be rendered by itself because it contains YAML front matter. We don't want this
because the index page will then have no content, because `index.html` does not have content --
it is a template. Instead, we want Jekyll to render content in `index.md` *using* 
`index.html` as template.

- `_includes/`

This directory contain reusable HTML parts that can be included in templates, such as
headers and footers.

- `_layouts/default.html`

This Liquid HTML layout is very simple, meant to be a template for other templates.
It includes the header and footer from `_includes/`. Since other templates will
be rendered using this template, the `{{ content }}` here is actuallly filled by
other HTML templates. 

- `_layouts/index.html`

This is a template for the index page. The front matter
indicates that it uses the `default.html` template.
In other words, it is wrapped by the default template, and whatever is here is put into
the `{{ content }}` section in `default.html`.

- `index.md`

Jekyll will render any file with YAML front matter, with no underscore
prepended to the filename.  Front matter is a block at
the top of the file enclosed with `---`. We create `index.md`, a markdown file,
with front matter signifying it is to be rendered with the index template,
via `layout: index`.
