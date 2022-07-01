package util

import "fmt"

func TreeQueryBuilder(objectid string) string {
	query := `WITH recursive ParentOfA
		AS
		(
		-- Anchor query
			SELECT 
				child_objectid, 
				(SELECT name from persons where objectid = relations.child_objectid) AS child_name, 
				parent_objectid,
				(SELECT name from persons where objectid = relations.parent_objectid) AS parent_name, 
				1 as level,
				(SELECT name from persons where objectid = relations.child_objectid) AS Hierarchy 
			FROM relations
			WHERE parent_objectid = '%s' OR child_objectid = '%s'
			UNION ALL
		-- Recursive query
			SELECT 
				r.child_objectid, 
				(SELECT name from persons where objectid = r.child_objectid) as child_name, 
				r.parent_objectid,
				(SELECT name from persons where objectid = r.parent_objectid) as parent_name, 
				M.level + 1 as level,
				M.Hierarchy || '->' || (SELECT name from persons where objectid = r.child_objectid) AS Hierarchy 
			FROM relations AS r
			JOIN ParentOfA AS M ON r.child_objectid = M. parent_objectid 
			
		), acendent AS (
			SELECT DISTINCT parent_objectid as objectid, parent_name as name FROM ParentOfA
			WHERE parent_objectid IS NOT NULL
			AND parent_objectid != '%s'
		), 
		ParentOfD
		AS
		(
		-- Anchor query
			SELECT 
				child_objectid, 
				(SELECT name from persons where objectid = relations.child_objectid) AS child_name, 
				parent_objectid,
				(SELECT name from persons where objectid = relations.parent_objectid) AS parent_name, 
				1 as level,
				(SELECT name from persons where objectid = relations.child_objectid) AS Hierarchy 
			FROM relations
			WHERE parent_objectid IS NULL
			UNION ALL
		-- Recursive query
			SELECT 
				r.child_objectid, 
				(SELECT name from persons where objectid = r.child_objectid) as child_name, 
				r.parent_objectid,
				(SELECT name from persons where objectid = r.parent_objectid) as parent_name, 
				M.level + 1 as level,
				M.Hierarchy || '->' || (SELECT name from persons where objectid = r.child_objectid) AS Hierarchy 
			FROM relations AS r
			JOIN ParentOfD AS M ON r.parent_objectid = M.child_objectid   
		), siblings AS (
			SELECT distinct D.child_objectid AS objectid, D.child_name AS name FROM ParentOfD as D, acendent AS U where D.parent_objectid = U.objectid
		),partialtree AS (
		SELECT * FROM siblings
		UNION
		SELECT * FROM acendent
		), decendents AS (
			SELECT distinct D.child_objectid AS objectid, D.child_name AS name FROM ParentOfD as D, partialtree AS U where D.parent_objectid = U.objectid
		), tree AS (
		SELECT * FROM partialtree
		UNION
		SELECT * FROM decendents
		)
		SELECT * from tree
		`
	return fmt.Sprintf(query, objectid, objectid, objectid)

}
