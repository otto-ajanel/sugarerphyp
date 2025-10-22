<?php

declare(strict_types=1);

namespace App\Infrastructure\Persistence\Eloquent\Models;


/**
 * @property int $id_category 
 * @property string $des_category 
 */
class Category extends TenantModel
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'categories';

    /**
     * The connection name for the model.
     */
    protected ?string $connection = 'pgsql';

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = [];
    protected string $primaryKey = 'id_category';

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['id_category' => 'integer'];
    public bool $timestamps = false;

}
