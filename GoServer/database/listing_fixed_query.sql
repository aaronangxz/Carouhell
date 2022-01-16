SELECT
    l.l_item_id,
    l.item_name,
    l.item_price,
    l.item_quantity,
    l.item_stock,
    l.item_description,
    l.item_location,
    l.item_status,
    l.item_category,
    l.l_seller_id,
    a.user_name AS seller_name,
    a.seller_type,
    l.listing_ctime,
    l.listing_mtime,
    COUNT(listing_reactions_tab.rt_item_id) as listing_likes
FROM
    acc_tab a,
    listing_tab l
    LEFT JOIN listing_reactions_tab ON l.l_item_id = listing_reactions_tab.rt_item_id
    AND listing_reactions_tab.reaction_type = ?
WHERE
    l.l_seller_id = a.a_user_id
    AND l.item_status = ?
GROUP BY
    l.l_item_id