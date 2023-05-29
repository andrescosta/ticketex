from setuptools import find_packages, setup
setup(
    name='tiklib',
    packages=find_packages(),
    version='0.1.0',
    description='Tiklib common library',
    author='Me',
    license='MIT',
    install_requires=[],
    setup_requires=['motor','pydantic','fastapi'],
)