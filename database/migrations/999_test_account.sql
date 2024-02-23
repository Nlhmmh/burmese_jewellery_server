INSERT INTO account_admins (account_admin_id, mail, password, account_admin_role, account_admin_status)
VALUES
  ('b106e919-ab73-4444-a86c-4b38ea22a0fe', 'admin@gmail.com', '$2a$10$On3UdcaBCZNYOwugRB9zuu5/AV8ywYVboiDqggpe0eaj8xK0Z1eTK', 'admin', 'active'),
  ('b106e919-ab73-4444-a86c-4b38ea22a0f1', 'staff@gmail.com', '$2a$10$On3UdcaBCZNYOwugRB9zuu5/AV8ywYVboiDqggpe0eaj8xK0Z1eTK', 'staff', 'active')
;

INSERT INTO categories (category_id, name, description, image_url)
VALUES
  ('b106e919-ab73-48f3-a86c-4b38ea22a0fe', 'Engagement Rings', 'Engagement Rings', 'https://media.tiffany.com/is/image/tiffanydm/2023_EngagementRing_BG_Hero-Desktop?$tile$&wid=2992&fmt=webp')
;

INSERT INTO gems (gem_id, name)
VALUES
  ('46e78b04-955d-4f0c-a751-dd1dcb2d7248', 'Diamond')
;

INSERT INTO materials (material_id, name)
VALUES
  ('c0979115-2bb3-4c50-a29d-64f3b8784a37', 'Gold')
;

INSERT INTO accounts (account_id, login_type, mail, password, phone, account_status)
VALUES
  ('b106e919-ab73-48f3-a86c-4b38ea22a888', 'email', 'testuser@gmail.com', '$2a$10$On3UdcaBCZNYOwugRB9zuu5/AV8ywYVboiDqggpe0eaj8xK0Z1eTK', '08077779999', 'active')
;

INSERT INTO account_profiles (account_id, first_name, last_name, birthday, gender)
VALUES
  ('b106e919-ab73-48f3-a86c-4b38ea22a888', 'John', 'Tester', '1990-11-01', 'male')
;

INSERT INTO faqs (faq_id, question, answer, is_active)
VALUES
  ('c0979115-2bb3-4c50-a29d-64f3b8784888', 'How can I make returns or exchanges?', 'Products are accepted for credit or exchange in the U.S. within 30 days of purchase. ', true),
  ('c0979115-2bb3-4c50-a29d-64f3b8784887', 'How do I return an online purchase?', 'Simply visit this page to print a return label.', false)
;

INSERT INTO jewelleries (jewellery_id, category_id, gem_id, material_id, name, description, price, image_url, is_published, quantity)
VALUES
  (
    'c0979115-2bb3-9999-a29d-64f3b8784888',
    'b106e919-ab73-48f3-a86c-4b38ea22a0fe',
    '46e78b04-955d-4f0c-a751-dd1dcb2d7248',
    'c0979115-2bb3-4c50-a29d-64f3b8784a37',
    'Wedding Forever',
    'Wedding Forever in Platinum with a Half-circle of Diamonds, 2.2 mm',
    '100',
    'https://media.tiffany.com/is/image/Tiffany/EcomItemL2/tiffany-foreverband-ring-16574635_1045538_ED_M.jpg?&op_usm=1.75,1.0,6.0&$cropN=0.1,0.1,0.8,0.8&defaultImage=NoImageAvailableInternal&&defaultImage=NoImageAvailableInternal&fmt=webp',
    true,
    100
  )
;
