<?php

namespace App\Application\Services;

use DateTimeImmutable;
use Lcobucci\JWT\Builder;
use Lcobucci\JWT\Configuration;
use Lcobucci\JWT\JwtFacade;
use Lcobucci\JWT\Signer\Hmac\Sha256;
use Lcobucci\JWT\Signer\Key\InMemory;
use Lcobucci\JWT\Validation\Constraint\SignedWith;
use Lcobucci\JWT\Validation\Constraint\StrictValidAt;
use Lcobucci\JWT\Validation\Constraint\ValidAt;
use Lcobucci\JWT\Validation\RequiredConstraintsViolated;

class Authservice
{

    public function encriptData($data)
    {
        $key = InMemory::base64Encoded(
            'hiG8DlOKvtih6AxlZn5XKImZ06yu8I3mkOzaJrEuW8yAv8Jnkw330uMt8AEqQ5LB'
        );

        $token = (new JwtFacade())->issue(
            new Sha256(),
            $key,
            static fn(
                Builder $builder,
                DateTimeImmutable $issuedAt
            ): Builder => $builder
                //->issuedBy('https://api.my-awesome-app.io')
                ->withClaim("data",$data)
                ->permittedFor('https://client-app.io')
                //->expiresAt($issuedAt->modify('+10 minutes'))
        );

        return $token->toString();
    }

    public function valideToken($tokenString){
         $key = InMemory::base64Encoded(
        'hiG8DlOKvtih6AxlZn5XKImZ06yu8I3mkOzaJrEuW8yAv8Jnkw330uMt8AEqQ5LB'
    );
    
    $config = Configuration::forSymmetricSigner(new Sha256(), $key);
    
    //try {
        $token = $config->parser()->parse($tokenString);
        
        // Constraints personalizadas
        $constraints = [
            new SignedWith($config->signer(), $config->signingKey())
        ];
        
        $config->validator()->assert($token, ...$constraints);
        return true;
      /*   
    } catch (RequiredConstraintsViolated $e) {
        // Token inválido
        return false;
    } catch (\Exception $e) {
        // Error en el formato del token
        return true;
    }
    */
    } 
}
