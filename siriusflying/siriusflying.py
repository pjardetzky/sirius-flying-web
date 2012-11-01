#
# $Id:$ 
#
import webapp2
from google.appengine.ext.webapp import template
import os

template_values = {
    'short_name': '',
    'long_name': '',
    }


class index(webapp2.RequestHandler):
    def get(self):
        path = os.path.join(os.path.dirname(__file__), 'templates/index.html')
        self.response.out.write(template.render(path, template_values))


class about(webapp2.RequestHandler):
    def get(self):
        path = os.path.join(os.path.dirname(__file__), 'templates/about.html')
        self.response.out.write(template.render(path, template_values))

class contact(webapp2.RequestHandler):
    def get(self):
        path = os.path.join(os.path.dirname(__file__), 'templates/contact.html')
        self.response.out.write(template.render(path, template_values))

class training(webapp2.RequestHandler):
    def get(self):
        path = os.path.join(os.path.dirname(__file__), 'templates/training.html')
        self.response.out.write(template.render(path, template_values))


class discovery(webapp2.RequestHandler):
    def get(self):
        path = os.path.join(os.path.dirname(__file__), 'templates/discovery.html')
        self.response.out.write(template.render(path, template_values))

class rental(webapp2.RequestHandler):
    def get(self):
        path = os.path.join(os.path.dirname(__file__), 'templates/rental.html')
        self.response.out.write(template.render(path, template_values))

class robots(webapp2.RequestHandler):
    def get(self):
        path = os.path.join(os.path.dirname(__file__), 'robots.txt')
        self.response.out.write(template.render(path, template_values))

app = webapp2.WSGIApplication(
                             [('/', index), 
                             ('/about', about),
                             ('/contact', contact),
                             ('/training', training),
                             ('/discovery', discovery),
                             ('/rental', rental),
                             ('/robots.txt', robots),
                             ],
                             debug=True)
