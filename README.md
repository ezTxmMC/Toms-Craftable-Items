# Tom's Craftable Items

## How To Use

### Create Item Recipe

### Shaped

```sh
./gen --dir data/minecraft/recipe --create SHAPED_CRAFTING minecraft:stick 16 'minecraft:oak_log[1,4]'
```

### Shapeless

```sh
./gen --dir data/minecraft/recipe --create SHAPELESS_CRAFTING minecraft:string 9 'minecraft:cobweb=1'
```

### Delete Item Recipe

```sh
./gen --dir data/minecraft/recipe --delete nether_star
```

## How To Build ZIP

```sh
./build.sh
```
