# Import modules
import requests
from YelpAPIKey import get_my_key

# Define a business id
business_id = '4AErMBEoNzbk7Qg45kKaQ'

# Define the API Key, Define the Endpoint, Define the Header
API_KEY = get_my_key()
ENDPOINT = 'https://api.yelp.com/v3/businesses/search'
HEADERS = {'Authorization': 'bearer %s' % API_KEY}

# Definte parameters
terms = 'restaurant'
location = 'Amherst'

PARAMETERS = {'terms': terms,
'limit': 10,
 'radius': 10000,
 'location': location}

 # Make a request to the yelp api
response = requests.get(url = ENDPOINT, params= PARAMETERS, headers= HEADERS)

# Convert JSON string to dictionary
business_data = response.json()

for biz in business_data['businesses']:
    print(biz['name'])