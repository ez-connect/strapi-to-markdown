# Strapi to Markdown

A simple application to fetch data from Strapi, then saving to markdown files. It support all Strapi content types, includes single type, collection type and media.

# Why

Strapi is a good CMS Headless for static sites but I'd like a really static site which:

- All data saved in a git repo
- Works fine without a server
- Compatible with both Gatsby & Hugo - my favorite tools

> Media in markdown content will not work without a server

# Usage

```
strapi-to-markdown -baseURL <baseURL> \
                    [-single <single_type_name>] \  
                    [-collection <collection_name>] \
                    [-body <body_field_name>] \
                    [-exclude <exclude_field_name>] \
                    -output <output_dir> \
                    [-static <static_dir>] \
                    -name <file_name_for_single_or_field_name_for_collection>

-baseURL string
    Strapi base url (default "http://localhost:1337")
-single string
    A single type name
-collection string
    A colletion type name
-body string
    Field name to write to markdown content
-exclude string
    Exclude field name
-output string
    Output directory
-static string
    Static directory (default "static")
-name string
    Output file name for single type or a field name of collection type   
```

# Example

Save the single type which named `nav` to `content/single/nav.md`

![](https://i.ibb.co/6vDy931/image.png)

```bash
strapi-to-markdown -single nav -output content/single -name nav.md
```

Save the single type which named `footer` to `content/single/footer.md`

```bash
strapi-to-markdown -single footer -output content/single -name footer.md
```

Save the collection type which name `post` to `content/posts`, use `slug` as file name, and `body` field to mardown body (other fields in frontmatter)

![](https://i.ibb.co/LRdfSJy/image.png)

```bash
strapi-to-markdown -collection post -body body -output content/posts -name slug
```

Save the collection type `tags` to `content/tags`, use `name` as file name and remove all referencens `posts` in data

![](https://i.ibb.co/YTBcgZ1/image.png)

```bash
strapi-to-markdown -collection tag -exclude posts -output content/tags -name name
```
