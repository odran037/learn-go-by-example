from lxml import html
import requests
import os
import sys

page = requests.get('https://gobyexample.com/')
tree = html.fromstring(page.content)
keywords = tree.xpath('//li//a/@href')

basedir = sys.argv[1] or 'learn-go-by-example'

def create_structure(dir):
    for index, keyword in enumerate(keywords):
        directory = '{}/{:02d}-{}'.format(dir, index, keyword)
        filename = '{:02d}-{}.go'.format(index, keyword)
        path  = '{}/{}'.format(directory, filename)
        os.makedirs(directory)
        if not os.path.exists(path):
            open(path, 'w').close()


if not os.path.exists(basedir):
    os.makedirs(basedir)
    create_structure(basedir)
else:
    print 'Directory already exists: {}'.format(basedir)
