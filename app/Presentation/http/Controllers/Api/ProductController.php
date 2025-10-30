<?php

declare(strict_types=1);

namespace App\Presentation\http\Controllers\Api;

use App\Application\Services\ProductService;
use Hyperf\HttpServer\Contract\RequestInterface;
use Hyperf\HttpServer\Contract\ResponseInterface;
use Psr\Http\Message\ResponseInterface as PsrRes;

class ProductController
{
    private ProductService $productService;
    private RequestInterface $req;
    private ResponseInterface $res;
    public function __construct(ProductService $productService, RequestInterface $request, ResponseInterface $response)
    {
        $this->productService = $productService;
        $this->req = $request;
        $this->res = $response;
    }
    public function createProduct() :PsrRes
    {
        
        $data = $this->req->all();
        $product = $this->productService->createProduct($data);
        return $this->res->json($product);

    }
    public function getAllProducts() :PsrRes
    {
        $products = $this->productService->getAllProducts();
        return $this->res->json($products);
    }
}
