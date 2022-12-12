import requests
from bs4 import BeautifulSoup

site = requests.get('http://www.dsc.ufcg.edu.br/~pet/jornal/maio2013/materias/tutoriais.html#:~:text=Em%20Python%2C%20existe%20um%20conceito,se%20comportar%C3%A1%20como%20uma%20fun%C3%A7%C3%A3o.')
print(site)

conteudo = BeautifulSoup(site, 'lxml')

print(conteudo.text)



