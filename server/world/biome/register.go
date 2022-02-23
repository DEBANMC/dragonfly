package biome

import "github.com/DEBANMC/dragonfly/server/world"

// init registers all biomes that can be used in a world.World.
func init() {
	world.RegisterBiome(Ocean{})
	world.RegisterBiome(LegacyFrozenOcean{})
	world.RegisterBiome(DeepOcean{})
	world.RegisterBiome(FrozenOcean{})
	world.RegisterBiome(DeepFrozenOcean{})
	world.RegisterBiome(ColdOcean{})
	world.RegisterBiome(DeepColdOcean{})
	world.RegisterBiome(LukewarmOcean{})
	world.RegisterBiome(DeepLukewarmOcean{})
	world.RegisterBiome(WarmOcean{})
	world.RegisterBiome(DeepWarmOcean{})
	world.RegisterBiome(River{})
	world.RegisterBiome(FrozenRiver{})
	world.RegisterBiome(Beach{})
	world.RegisterBiome(StonyShore{})
	world.RegisterBiome(SnowyBeach{})
	world.RegisterBiome(Forest{})
	world.RegisterBiome(WoodedHills{})
	world.RegisterBiome(FlowerForest{})
	world.RegisterBiome(BirchForest{})
	world.RegisterBiome(BirchForestHills{})
	world.RegisterBiome(OldGrowthBirchForest{})
	world.RegisterBiome(TallBirchHills{})
	world.RegisterBiome(DarkForest{})
	world.RegisterBiome(DarkForestHills{})
	world.RegisterBiome(Jungle{})
	world.RegisterBiome(JungleHills{})
	world.RegisterBiome(ModifiedJungle{})
	world.RegisterBiome(JungleEdge{})
	world.RegisterBiome(ModifiedJungleEdge{})
	world.RegisterBiome(BambooJungle{})
	world.RegisterBiome(BambooJungleHills{})
	world.RegisterBiome(Taiga{})
	world.RegisterBiome(TaigaHills{})
	world.RegisterBiome(TaigaMountains{})
	world.RegisterBiome(SnowyTaiga{})
	world.RegisterBiome(SnowyTaigaHills{})
	world.RegisterBiome(SnowyTaigaMountains{})
	world.RegisterBiome(OldGrowthPineTaiga{})
	world.RegisterBiome(GiantTreeTaigaHills{})
	world.RegisterBiome(OldGrowthSpruceTaiga{})
	world.RegisterBiome(GiantSpruceTaigaHills{})
	world.RegisterBiome(MushroomFields{})
	world.RegisterBiome(MushroomFieldShore{})
	world.RegisterBiome(Swamp{})
	world.RegisterBiome(SwampHills{})
	world.RegisterBiome(Savanna{})
	world.RegisterBiome(SavannaPlateau{})
	world.RegisterBiome(WindsweptSavanna{})
	world.RegisterBiome(ShatteredSavannaPlateau{})
	world.RegisterBiome(Plains{})
	world.RegisterBiome(SunflowerPlains{})
	world.RegisterBiome(Desert{})
	world.RegisterBiome(DesertHills{})
	world.RegisterBiome(DesertLakes{})
	world.RegisterBiome(SnowyPlains{})
	world.RegisterBiome(SnowyMountains{})
	world.RegisterBiome(IceSpikes{})
	world.RegisterBiome(GravellyMountainsPlus{})
	world.RegisterBiome(MountainEdge{})
	world.RegisterBiome(Badlands{})
	world.RegisterBiome(BadlandsPlateau{})
	world.RegisterBiome(ModifiedBadlandsPlateau{})
	world.RegisterBiome(WoodedBadlandsPlateau{})
	world.RegisterBiome(ModifiedWoodedBadlandsPlateau{})
	world.RegisterBiome(ErodedBadlands{})
	world.RegisterBiome(Meadow{})
	world.RegisterBiome(Grove{})
	world.RegisterBiome(SnowySlopes{})
	world.RegisterBiome(JaggedPeaks{})
	world.RegisterBiome(FrozenPeaks{})
	world.RegisterBiome(StonyPeaks{})
	world.RegisterBiome(LushCaves{})
	world.RegisterBiome(DripstoneCaves{})
	world.RegisterBiome(NetherWastes{})
	world.RegisterBiome(CrimsonForest{})
	world.RegisterBiome(WarpedForest{})
	world.RegisterBiome(SoulSandValley{})
	world.RegisterBiome(BasaltDeltas{})
	world.RegisterBiome(End{})
}
