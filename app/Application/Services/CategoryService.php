<?php
namespace App\Application\Services;

use App\Infrastructure\Persistence\Eloquent\Models\Category;

class CategoryService{
    
    public function getAllCategories(): array
    {
       return Category::all()->toArray();
    }
    public function createCategory(array $data): array
    {
        $newCategory = new Category();
        $newCategory->des_category = $data['newCategory'];
        $newCategory->save();

        return $newCategory->toArray();
    }
}