from setuptools import setup, find_packages
from pathlib import Path

# Read README.md for PyPI long_description
this_dir = Path(__file__).parent
long_description = (this_dir / "README.md").read_text(encoding="utf-8")

setup(
    name="discardapi-demo",
    version="0.1.0",
    description="Python demo client for Discard API",
    long_description=long_description,
    long_description_content_type="text/markdown",  # Important!
    author="Qasim Ali",
    author_email="discardapi@gmail.com",
    url="https://github.com/GlobalTechInfo/discard-api",
    packages=find_packages(),
    install_requires=["requests"],
    python_requires=">=3.7",
)
