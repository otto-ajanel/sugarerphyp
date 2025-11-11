<?php  
namespace App\Presentation\http\Controllers\Api;

use App\Application\Services\AtributeService;
use Hyperf\HttpServer\Contract\RequestInterface;
use Hyperf\HttpServer\Contract\ResponseInterface;
use Psr\Http\Message\ResponseInterface as PsrRes;


class AtributeController
{
    private AtributeService $atributeService;
    private RequestInterface $req;
    private ResponseInterface $res;
    public function __construct(AtributeService $atributeService, RequestInterface $request, ResponseInterface $response)
    {
        $this->atributeService = $atributeService;
        $this->req = $request;
        $this->res = $response;
    }

    public function createAtribute(): PsrRes
    {
        $data = $this->req->all();
        $atribute = $this->atributeService->createAtribute($data);
        return $this->res->json($atribute);
    }

    public function getAllAtributes(): PsrRes
    {
        return $this->res->json($this->atributeService->getAllAtributes());
    }
}
