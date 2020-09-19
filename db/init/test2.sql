SELECT
		   id user_id, name, ranking, high_score score
	  	 FROM
		   (
		     SELECT
			   CASE
			     WHEN @p=high_score
				   THEN @c
				 ELSE @c:=@c+@s
			   END AS ranking,
			   CASE
				 WHEN @p=high_score
				   THEN @s:=@s+1
				 ELSE @s:=1
			   END AS dummy1,
			   id, name, high_score, @p:=high_score
		     FROM
			   (
				 SELECT
				   @c:=1, @s:=0
			   ) AS dummy2,
			   user
		     ORDER BY
			   high_score
		     DESC
		   )
	  	 AS tmp
	     LIMIT 0, 10;