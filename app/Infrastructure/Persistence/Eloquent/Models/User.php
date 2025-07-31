<?php

declare(strict_types=1);

namespace App\Infrastructure\Persistence\Eloquent\Models;

use Hyperf\DbConnection\Model\Model;

/**
 * @property int $id_user 
 * @property string $username 
 * @property string $email 
 * @property string $password 
 * @property int $id_usertype 
 * @property \Carbon\Carbon $created_at 
 * @property boolean $active 
 * @property int $id_tenance 
 */
class User extends Model
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'users';

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
    protected array $casts = ['id_user' => 'integer', 'id_usertype' => 'integer', 'created_at' => 'datetime', 'active' => 'boolean', 'id_tenance' => 'integer'];
}
