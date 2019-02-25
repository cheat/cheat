def is_tool(name):
    """Check whether `name` is on PATH."""

    from distutils.spawn import find_executable

    return find_executable(name) is not None

def outPager(sheet_content,colorize):
    """Runs pydoc.pager(sheet_content). If less is on PATH, colorizes output"""
    if is_tool('less'):
        from pydoc import pipepager
        pipepager(colorize.syntax(sheet_content),'less -R')
    else:
        from pydoc import pager
        pager(sheet_content)