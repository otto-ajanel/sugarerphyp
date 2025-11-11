<?php

declare(strict_types=1);

namespace App\Infrastructure\Persistence\Eloquent\Models;

use Hyperf\DbConnection\Model\Model;

/**
 * @property int $atribute_id 
 * @property int $category_id 
 * @property string $atribute_des 
 * @property string $atribute_typedata 
 */
class Atribute extends TenantModel
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'atribute';

    /**
     * The connection name for the model.
     */
    protected ?string $connection = 'pgsql';

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = [];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['atribute_id' => 'integer', 'category_id' => 'integer'];

    public bool $timestamps = false;

    protected string $primaryKey = 'atribute_id';
}
