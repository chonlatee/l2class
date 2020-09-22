package cmd

import (
	"fmt"
	"strings"

	"github.com/chonlatee/l2class/internal/lineageclass"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of all class",
	RunE: func(cmd *cobra.Command, args []string) error {

		conditions := []map[string]string{}

		race, err := cmd.Flags().GetString("race")

		if err != nil {
			return err
		}

		if len(race) > 0 {
			cond := map[string]string{
				"race": race,
			}

			conditions = append(conditions, cond)
		}

		starterClass, err := cmd.Flags().GetString("starterClass")

		if len(starterClass) > 0 {
			cond := map[string]string{
				"starterClass": starterClass,
			}

			conditions = append(conditions, cond)
		}

		firstClass, err := cmd.Flags().GetString("firstClass")

		if len(firstClass) > 0 {
			cond := map[string]string{
				"firstClass": firstClass,
			}

			conditions = append(conditions, cond)
		}

		secondClass, err := cmd.Flags().GetString("secondClass")

		if len(secondClass) > 0 {
			cond := map[string]string{
				"secondClass": secondClass,
			}

			conditions = append(conditions, cond)
		}

		thirdClass, err := cmd.Flags().GetString("thirdClass")

		if len(secondClass) > 0 {
			cond := map[string]string{
				"thirdClass": thirdClass,
			}

			conditions = append(conditions, cond)
		}

		return list(conditions)
	},
}

func init() {
	listCmd.Flags().String("race", "", "-race=human")
	listCmd.Flags().String("starterClass", "", "-starterClass=fighter")
	listCmd.Flags().String("firstClass", "", "-firstClass=knight")
	listCmd.Flags().String("secondClass", "", "-secondClass=paladin")
	listCmd.Flags().String("thirdClass", "", "-thirdClass=phoenix Knight")
	rootCmd.AddCommand(listCmd)
}

func list(conditions []map[string]string) error {

	listClass, err := lineageclass.LoadClass()

	if err != nil {
		return err
	}

	list := filterByCondition(listClass, conditions)

	for _, v := range list {
		fmt.Println()
		fmt.Printf("[ %s -> %s -> %s -> %s -> %s ]\n", v.Race, v.StarterClassName, v.FirstClassName, v.SecondClassName, v.ThirdClassName)
		fmt.Printf("Description: %s\n", v.Description)
		fmt.Printf("Pros: %s\n", v.Pros)
		fmt.Printf("Cons: %s\n", v.Cons)
	}

	return nil
}

// Bad code need to improve
func filterByCondition(list []*lineageclass.Class, conditions []map[string]string) []*lineageclass.Class {
	l := []*lineageclass.Class{}
	_, valRace := containCondition("race", conditions)
	_, valStarterClass := containCondition("starterClass", conditions)
	_, valFirstClass := containCondition("firstClass", conditions)
	_, valSecondClass := containCondition("secondClass", conditions)
	_, valThirdClass := containCondition("thirdClass", conditions)

	l = list

	if len(valRace) > 0 {
		l = filterRace(valRace, l)
	}

	if len(valStarterClass) > 0 {
		l = filterStarterClass(valStarterClass, l)
	}

	if len(valFirstClass) > 0 {
		l = filterFirstClass(valFirstClass, l)
	}

	if len(valSecondClass) > 0 {
		l = filterSecondClass(valSecondClass, l)
	}

	if len(valThirdClass) > 0 {
		l = filterThirdClass(valThirdClass, l)
	}

	return l
}

func containCondition(cond string, conditions []map[string]string) (bool, string) {

	for _, v := range conditions {
		if val, ok := v[cond]; ok {
			return true, val
		}
	}

	return false, ""
}

// All filter function is bad code ned to imporve
func filterRace(race string, list []*lineageclass.Class) []*lineageclass.Class {
	l := []*lineageclass.Class{}

	for _, v := range list {
		if strings.ToLower(v.Race) == race {
			l = append(l, v)
		}
	}

	if len(l) > 0 {
		return l
	}

	return list
}

func filterStarterClass(class string, list []*lineageclass.Class) []*lineageclass.Class {
	l := []*lineageclass.Class{}

	for _, v := range list {
		if strings.ToLower(v.StarterClassName) == class {
			l = append(l, v)
		}
	}

	if len(l) > 0 {
		return l
	}

	return list
}

func filterFirstClass(class string, list []*lineageclass.Class) []*lineageclass.Class {
	l := []*lineageclass.Class{}

	for _, v := range list {
		if strings.ToLower(v.FirstClassName) == class {
			l = append(l, v)
		}
	}

	if len(l) > 0 {
		return l
	}

	return list
}

func filterSecondClass(class string, list []*lineageclass.Class) []*lineageclass.Class {
	l := []*lineageclass.Class{}

	for _, v := range list {
		if strings.ToLower(v.SecondClassName) == class {
			l = append(l, v)
		}
	}

	if len(l) > 0 {
		return l
	}

	return list
}

func filterThirdClass(class string, list []*lineageclass.Class) []*lineageclass.Class {
	l := []*lineageclass.Class{}

	for _, v := range list {
		if strings.ToLower(v.ThirdClassName) == class {
			l = append(l, v)
		}
	}

	if len(l) > 0 {
		return l
	}

	return list
}
