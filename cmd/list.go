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

type listBuilder struct {
	list []*lineageclass.Class
}

// a lot of duplicate code need to imporve
func (l *listBuilder) filterByRace(race string) *listBuilder {
	newList := []*lineageclass.Class{}

	if len(race) <= 0 {
		return l
	}

	for _, v := range l.list {
		if strings.ToLower(v.Race) == strings.ToLower(race) {
			newList = append(newList, v)
		}
	}

	l.list = newList

	return l

}

func (l *listBuilder) filterByStarterClass(starterClass string) *listBuilder {
	newList := []*lineageclass.Class{}

	if len(starterClass) <= 0 {
		return l
	}

	for _, v := range l.list {
		if strings.ToLower(v.StarterClassName) == strings.ToLower(starterClass) {
			newList = append(newList, v)
		}
	}

	l.list = newList

	return l
}

func (l *listBuilder) filterByFirstClass(firstClass string) *listBuilder {
	newList := []*lineageclass.Class{}

	if len(firstClass) <= 0 {
		return l
	}

	for _, v := range l.list {
		if strings.ToLower(v.FirstClassName) == strings.ToLower(firstClass) {
			newList = append(newList, v)
		}
	}

	l.list = newList

	return l
}

func (l *listBuilder) filterBySecondClass(secondClass string) *listBuilder {
	newList := []*lineageclass.Class{}

	if len(secondClass) <= 0 {
		return l
	}

	for _, v := range l.list {
		if strings.ToLower(v.SecondClassName) == strings.ToLower(secondClass) {
			newList = append(newList, v)
		}
	}

	l.list = newList

	return l
}

func (l *listBuilder) filterByThirdClass(thirdClass string) *listBuilder {
	newList := []*lineageclass.Class{}

	if len(thirdClass) <= 0 {
		return l
	}

	for _, v := range l.list {
		if strings.ToLower(v.ThirdClassName) == strings.ToLower(thirdClass) {
			newList = append(newList, v)
		}
	}

	l.list = newList

	return l
}

// Bad code need to improve
func filterByCondition(list []*lineageclass.Class, conditions []map[string]string) []*lineageclass.Class {
	_, race := containCondition("race", conditions)
	_, starterClass := containCondition("starterClass", conditions)
	_, firstClass := containCondition("firstClass", conditions)
	_, secondClass := containCondition("secondClass", conditions)
	_, thirdClass := containCondition("thirdClass", conditions)

	bList := &listBuilder{
		list: list,
	}

	bList.filterByRace(race).
		filterByStarterClass(starterClass).
		filterByFirstClass(firstClass).
		filterBySecondClass(secondClass).
		filterByThirdClass(thirdClass)

	return bList.list
}

func containCondition(cond string, conditions []map[string]string) (bool, string) {

	for _, v := range conditions {
		if val, ok := v[cond]; ok {
			return true, val
		}
	}

	return false, ""
}
