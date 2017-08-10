package pve

import (
	"encoding/json"
	"log"
	"math"
	"slg/common"
	"slg/model"
	"time"
)

func (this *Pve) attack(param interface{}, extra interface{}) (output []byte, errno int) {

	if paramData, ok := param.(map[string]interface{}); ok {

		if myData, ok := extra.(*common.UserData); ok {

			cityId64, _ := paramData["cityId"].(uint64)
			cityId := int(cityId64)

			corpsId64, _ := paramData["corpsId"].(uint64)
			corpsId := int(corpsId64)

			// 出发点坐标
			spos64, _ := paramData["spos"].(uint64)
			spos := int(spos64)

			// 到达点坐标
			epos64, _ := paramData["epos"].(uint64)
			epos := int(epos64)

			attackInfo := &AttackInfoStruct{}
			attackInfo.CityId = cityId
			attackInfo.Atime = time.Now().Unix()
			attackInfo.CorpsId = corpsId
			attackInfo.Spos = spos
			attackInfo.Epos = epos

			atkByte, err := json.Marshal(attackInfo)
			if err != nil {
				return []byte{}, common.CODE_ERR_UNKNOW
			}

			pveMapObj := model.NewPveMap()

			res := pveMapObj.Atk(myData.Uid, corpsId, string(atkByte))

			if !res {
				return []byte{}, common.CODE_ERR_UNKNOW
			}

			atkMP := &AttackInfoMpStruct{}
			atkMP.Id = CODE_ATTACK
			atkMP.Time = time.Now().Unix()
			atkMP.CityId = cityId
			atkMP.Atime = time.Now().Unix()
			atkMP.CorpsId = corpsId
			atkMP.Spos = spos
			atkMP.Epos = epos

			// 读取部队信息
			armyObj := model.NewArmy()
			log.Println("army obj:", myData.Uid, cityId, corpsId)

			armyData := armyObj.GetArmyByCid(myData.Uid, cityId, corpsId)

			// 读取被攻击方的部队信息
			x := int(math.Floor(float64(epos64 / 100)))
			y := epos % 100
			var landType int = 0

			for _, v := range common.PveMapConfs {
				if v.X == x && v.Y == y {
					landType = v.Type
				}
			}

			if landType != 0 {
				pveCorpId := 0
				//pveCorpNum := 0
				//				pveHp := 0

				for _, v := range common.GroundTypeConfs {
					if v.Id == landType {
						pveCorpId = v.PveCorpsId
						//pveCorpNum = v.PveCorpsNum
						//pveHp = v.PveHp
					}
				}

				// 触发战斗定时器
				//go func(myArmy map[int]int, pveCorpId int, pveCorpNum int, pveHp int) {

				// 生成敌军部队
				enemy := make(map[int]int)

				for _, v := range common.NpcCorpsConfs {
					if v.CorpsId == pveCorpId {
						if v.Aid != 0 {
							enemy[v.Aid] = v.Num
						}
						if v.Aid1 != 0 {
							enemy[v.Aid1] = v.Num1
						}
						if v.Aid2 != 0 {
							enemy[v.Aid2] = v.Num2
						}
						if v.Aid3 != 0 {
							enemy[v.Aid3] = v.Num3
						}
						if v.Aid4 != 0 {
							enemy[v.Aid4] = v.Num4
						}
					}
				}

				// 编排站位
				myFormattion := make([]*FightFormation, 0, 16)
				atkMP.MyArmy = make([]*FightFormation, 0, 16)
				pos := 101
				for k, v := range armyData {
					armyFormation := &FightFormation{}
					armyFormation.Aid = k
					armyFormation.Num = v
					armyFormation.Pos = pos
					myFormattion = append(myFormattion, armyFormation)

					armyOutInfo := &FightFormation{}
					armyOutInfo.Aid = k
					armyOutInfo.Num = v
					armyOutInfo.Pos = pos
					atkMP.MyArmy = append(atkMP.MyArmy, armyOutInfo)

					pos += 1

				}

				enemyFormattion := make([]*FightFormation, 0, 16)
				atkMP.EnemyArmy = make([]*FightFormation, 0, 16)

				pos = 1001
				for k, v := range enemy {
					armyFormation := &FightFormation{}
					armyFormation.Aid = k
					armyFormation.Num = v
					armyFormation.Pos = pos
					enemyFormattion = append(enemyFormattion, armyFormation)

					armyOutInfo := &FightFormation{}
					armyOutInfo.Aid = k
					armyOutInfo.Num = v
					armyOutInfo.Pos = pos
					atkMP.EnemyArmy = append(atkMP.EnemyArmy, armyOutInfo)

					pos += 1
				}

				// 开战
				attacker := 1 // 攻击方 1 本方 2 敌人

				myCurrent := 0    // 本方攻击游标
				enemyCurrent := 0 // 敌方攻击游标
				fightRecord := make([]*FightRecord, 0, 1024)
				var myAtkVal, enemyAtkVal, def float64
				for {
					if attacker == 1 {
						log.Println("I attack")

						attacker = 0
						atkInfo := &FightRecord{}
						atkInfo.Atk = myFormattion[myCurrent].Pos

						if myFormattion[myCurrent].Num > 0 {
							myAtkVal = float64(common.ArmyConfs[myFormattion[myCurrent].Aid].Atk * myFormattion[myCurrent].Num)

							// 获取可攻击的敌方部队
							for k, v := range enemyFormattion {

								if v.Num > 0 {
									// 获取对方的防御力
									def = float64(common.ArmyConfs[v.Aid].Def * v.Num)
									hurt := int(math.Ceil(float64(v.Num) * (myAtkVal / def)))
									log.Println("enemy army:", v.Aid, v.Num, myAtkVal, def, hurt)

									enemyFormattion[k].Num -= hurt
									hurtRecord := &HurtRecord{}
									hurtRecord.Hurt = hurt
									hurtRecord.Target = v.Pos
									atkInfo.Targets = append(atkInfo.Targets, hurtRecord)
									fightRecord = append(fightRecord, atkInfo)
									attacker = 2
									break
								}
							}
						}
						myCurrent++
						myCurrent = myCurrent % len(myFormattion)

						if attacker == 0 {
							atkMP.Win = 1
							break
						}
					} else {
						attacker = 0
						atkInfo := &FightRecord{}
						atkInfo.Atk = enemyFormattion[enemyCurrent].Pos

						if enemyFormattion[enemyCurrent].Num > 0 {
							enemyAtkVal = float64(common.ArmyConfs[enemyFormattion[enemyCurrent].Aid].Atk * enemyFormattion[enemyCurrent].Num)

							// 获取可攻击的我方部队
							for k, v := range myFormattion {
								if v.Num > 0 {

									// 获取我方的防御力
									def = float64(common.ArmyConfs[v.Aid].Def * v.Num)
									hurt := int(math.Ceil(float64(v.Num) * (enemyAtkVal / def)))
									log.Println("my army:", v.Aid, v.Num, enemyAtkVal, def, hurt)

									myFormattion[k].Num -= hurt
									hurtRecord := &HurtRecord{}
									hurtRecord.Hurt = hurt
									hurtRecord.Target = v.Pos
									atkInfo.Targets = append(atkInfo.Targets, hurtRecord)
									fightRecord = append(fightRecord, atkInfo)
									attacker = 1
									break
								}
							}
						}
						enemyCurrent++
						enemyCurrent = enemyCurrent % len(enemyFormattion)

						if attacker == 0 {
							atkMP.Win = 0
							break
						}
					}
				}

				atkMP.FightRecord = fightRecord

				//ticker := time.NewTicker(10 * time.Second)

				//select {
				//case <-ticker.C:
				// 计算战斗结果
				//}
				//}(armyData, pveCorpId, pveCorpNum, pveHp)
			}

			serialize, err := common.Serialize(atkMP)

			if err != nil {
				return []byte{}, common.CODE_ERR_UNKNOW
			}

			return serialize, 0

		}

	}
	return []byte{}, 0
}

func computeFight() {

}
