# Create new environment
virtualenv /path/to/project/env_name

# Create new environment and inherit already installed Python libraries
virtualenv --system-site-package /path/to/project/env_name

# Create new environment with a given Python interpreter
virtualenv /path/to/project/env_name -p /usr/bin/python/3.4

# Activate environnment
source /path/to/project/env_name/bin/activate

# Quit environment
deactivate


# virtualenvwrapper (wrapper for virtualenv)
# installation
pip install --user virtualenvwrapper
# configuration
# add in ~/.bashrc or similar
export WORKON_HOME=~/.virtualenvs
mkdir -p $WORKON_HOME
source ~/.local/bin/virtualenvwrapper.sh

# Create new environmment (with virtualenvwrapper)
mkvirtualenv env_name
# new environmment is stored in ~/.virtualenvs

# Activate environmment (with virtualenvwrapper)
workon env_name

# Quit environmment (with virtualenvwrapper)
deactivate

# Delete environmment (with virtualenvwrapper)
rmvirtualenv env_name

