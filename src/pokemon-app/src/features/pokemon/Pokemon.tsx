import { Pokemon as PokemonInterface } from "../../app/models/Pokemon";

export function Pokemon({ Id, Name, Image }: PokemonInterface) {


    return (<div>
        {Name ?
            <div className="pokemon"><div className="title">{Name}</div><div className="descrtiption">Pokem ipsum dolor sit amet Azumarill Bidoof Gible Hippowdon consectetur adipisicing elit Regirock. Body Slam Helix Fossil Yanma Onix Abomasnow Machop Pupitar. Hoenn Poison Sting Kricketot Bayleef Rising Badge Abra Sandslash. Duis aute irure dolor in reprehenderit in voluptate Silver Tangela Lopunny Mantine Cleffa Mantine. Kanto Timburr Claydol Reuniclus consectetur adipisicing elit Pansear Dialga.</div><div className="image"><img src={Image} alt={Name} /></div></div>
            : <div>no hay nada para ti </div>}
    </div>
    );
}