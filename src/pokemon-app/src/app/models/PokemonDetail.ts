export interface PokemonDetail {
  abilities: ability[];
  base_experience: number;
  forms: form[];
  game_indices: game_index[];
  height: number;
  id: number;
  is_default: boolean;
  location_area_encounters: string;
  moves: move[];
  name: string;
  order: number;
  weight: number;
}

export interface ability {
  ability: baseInterface;
  is_hidden: boolean;
  slot: number;
}

interface form extends baseInterface {}

interface game_index {
  game_index: number;
  version: game_version;
}

interface game_version extends baseInterface {}

interface move {
  move: baseInterface;
  version_group_details: version_group_detail[];
}

interface version_group_detail {
  level_learned_at: number;
  move_learn_method: baseInterface;
  version_group: baseInterface;
}

interface baseInterface {
  name: string;
  url: string;
}
