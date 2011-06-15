#
# $Id: siriusflying.py,v a9bff61f36d3 2011/06/15 03:19:28 pjardetzky $ 
#
# @author         $Author$
# @version        $Rev$
# @lastrevision   $Date$
#
from google.appengine.ext import webapp
from google.appengine.ext.webapp.util import run_wsgi_app
from google.appengine.ext.webapp import template
import os

template_values = {
    'short_name': '',
    'long_name': '',
    }


class index(webapp.RequestHandler):
    def get(self):
        path = os.path.join(os.path.dirname(__file__), 'templates/index.html')
        self.response.out.write(template.render(path, template_values))


class about(webapp.RequestHandler):
    def get(self):
        path = os.path.join(os.path.dirname(__file__), 'templates/about.html')
        self.response.out.write(template.render(path, template_values))

class contact(webapp.RequestHandler):
    def get(self):
        path = os.path.join(os.path.dirname(__file__), 'templates/contact.html')
        self.response.out.write(template.render(path, template_values))

class training(webapp.RequestHandler):
    def get(self):
        path = os.path.join(os.path.dirname(__file__), 'templates/training.html')
        self.response.out.write(template.render(path, template_values))


class discovery(webapp.RequestHandler):
    def get(self):
        path = os.path.join(os.path.dirname(__file__), 'templates/discovery.html')
        self.response.out.write(template.render(path, template_values))

class rental(webapp.RequestHandler):
    def get(self):
        path = os.path.join(os.path.dirname(__file__), 'templates/rental.html')
        self.response.out.write(template.render(path, template_values))

class robots(webapp.RequestHandler):
    def get(self):
        path = os.path.join(os.path.dirname(__file__), 'robots.txt')
        self.response.out.write(template.render(path, template_values))

application = webapp.WSGIApplication(
                                     [('/', index), 
                                      ('/about', about),
                                      ('/contact', contact),
                                      ('/training', training),
                                      ('/discovery', discovery),
                                      ('/rental', rental),
                                      ('/robots.txt', robots),
                                      ],
                                     debug=True)

def main():
    run_wsgi_app(application)

################################################################################
if __name__ == "__main__":
    main()
