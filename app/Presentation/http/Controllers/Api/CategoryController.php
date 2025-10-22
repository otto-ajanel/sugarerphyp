<?php
declare(strict_types=1);
namespace App\Presentation\http\Controllers\Api;

use App\Application\Services\CategoryService;
use Hyperf\HttpServer\Contract\RequestInterface;
use Hyperf\HttpServer\Contract\ResponseInterface;
use Psr\Http\Message\ResponseInterface as PsrRes;

class CategoryController
{
    private $categoryService;
    private RequestInterface $req;
    private ResponseInterface $res;


    public function __construct(
        CategoryService $categoryService,
        RequestInterface $request,
        ResponseInterface $response
        )
    {
        $this->categoryService = $categoryService;
        $this->req = $request;
        $this->res = $response;
    }

    public function  index() :PsrRes
    {
        try {
            $reqData = $this->req->all();


            $data=$this->categoryService->getAllCategories($reqData);
            return $this->res->json($data);

        } catch (\Throwable $th) {
            return $this->res->json(['error' => $th->getMessage()])->withStatus(501);
        }
    }

    public function store() :PsrRes
    {
        try {
            $reqData = $this->req->all();


            $data=$this->categoryService->createCategory($reqData);
            return $this->res->json($data);

        } catch (\Throwable $th) {
            return $this->res->json(['error' => $th->getMessage()])->withStatus(501);
        }
    }

    
}