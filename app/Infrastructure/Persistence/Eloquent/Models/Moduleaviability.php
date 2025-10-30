<?php

declare(strict_types=1);

namespace App\Infrastructure\Persistence\Eloquent\Models;

use Hyperf\DbConnection\Model\Model;

/**
 * @property int $modavia_id 
 * @property int $module_id 
 * @property int $product_id 
 */
class Moduleaviability extends TenantModel
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'moduleaviability';

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = [
        'module_id',
        'product_id'
    ];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['modavia_id' => 'integer', 'module_id' => 'integer', 'product_id' => 'integer'];
    protected string $primaryKey = 'modavia_id';
    public bool $timestamps = false;    
}
