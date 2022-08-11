#!/usr/bin/env python3
import subprocess
import json

print(subprocess.run(['detect-secrets', 'scan', '--update', '.secrets.baseline']))

found_secrets = []

with open('.secrets.baseline', 'r') as f:
    baseline = json.loads(f.read())
    for file, secrets in baseline['results'].items():
        for secret in secrets:
            if secret.get('is_secret', True):
                found_secrets.append((file, secret))

if found_secrets:
    print('Secrets were found in the source code!')
    print('If these contain false positives, they can be marked as such with the `detect-secrets audit .secrets.baseline` command and committing the updated baseline file into the application repo.')
    print('Read more about the tool at https://github.com/ibm/detect-secrets#about\n\n')
    print('FOUND SECRETS:')
    for secret in found_secrets:
        print('File: ' + secret[0] + ' Line: ' + str(secret[1]['line_number']) + ' Type: ' + secret[1]['type'])
    print('failure')
    exit(1)
else:
    print('NO SECRETS FOUND')
    print('success')

