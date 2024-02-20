INSERT INTO account_admins (mail, password, account_admin_role, account_admin_status)
VALUES
  ('admin@gmail.com', '$2a$10$On3UdcaBCZNYOwugRB9zuu5/AV8ywYVboiDqggpe0eaj8xK0Z1eTK', 'admin', 'active')
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
