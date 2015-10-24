# Party Template Handler for Go Web development

Web Template handler that supports partial template.

Put your template files in ./template directory;

Put your partial template files in ./template/partial directory

In your template file, use:

 {{include "partial_filename"}}

 to include the partial template.


Create a template handler with:
party.New(filename string, data map[string]interface{}, development bool)

If the 'development' argument is true, use dyanamic template that facilitate development-time flexibility: change the template file and refresh browser to see the result. Otherwise, use static template that maximize the run-time speed.

That's it, enjoy!
