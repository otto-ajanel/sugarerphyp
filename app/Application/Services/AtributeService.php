<?php   
namespace App\Application\Services;

use App\Infrastructure\Persistence\Eloquent\Models\Atribute;

class AtributeService
{
    private Atribute $atributeModel;
    public function __construct(Atribute $atributeModel)
    {
        $this->atributeModel = $atributeModel;
    }

    public function createAtribute(array $data): Atribute
    {

        $atribute = new Atribute();
        $atribute->category_id = $data['categoryId'];
        $atribute->atribute_des = $data['atributeDes'];
        $atribute->atribute_typedata = $data['atributeTypedata'];

        $atribute->save();
        return $atribute;
    }
    public function getAllAtributes(): array
    {
        return $this->atributeModel->all()->toArray();
    }
}