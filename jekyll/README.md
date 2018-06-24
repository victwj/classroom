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

- [Bundler](https://bundler.io) is another important Ruby gem. It is a tool to help manage project
environments, and makes it easy to manage dependencies (i.e. other gems, their dependencies, their versions).

- Jekyll transforms text. Without a tool like Jekyll, devs need to write content in pure HTML. With Jekyll, we can
break this down to a more maintainable process. We first make templates for our website, in Liquid HTML, to define
what things should look like. Then we write our content in Markdown. Jekyll will connect the two together and put the
content in the templates.

- To get started:

```bash
mkdir website
cd website

# Generate a Gemfile in current directory
# Gemfile is how Bundler knows what the required dependencies are, and how to install them
bundle init

# Add the jeykll gem to Gemfile
bundle add jekyll

# Install the dependencies specified by Gemfile
bundle install
```
